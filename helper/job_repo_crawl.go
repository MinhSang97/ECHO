package helper

import (
	"app/banana"
	"app/log"
	"app/model"
	"app/repository"
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func CrawlRepo(githubRepo repository.GithubRepo) {
	c := colly.NewCollector()

	repos := make([]model.GithubRepo, 0, 30)
	c.OnHTML("article.Box-row", func(e *colly.HTMLElement) {
		var githubRepo model.GithubRepo
		// repair before crawler no data
		githubRepo.Name = strings.Join(strings.Fields(e.DOM.Find("h2").Text()), " ")
		fmt.Println(githubRepo.Name)
		if githubRepo.Name == "" {
			fmt.Println("error: khong co data")
		}

		githubRepo.Description = e.ChildText("p.col-9")

		bgColor := e.ChildAttr(".repo-language-color", "style")
		re := regexp.MustCompile("#[a-zA-Z0-9_]+")
		match := re.FindStringSubmatch(bgColor)
		if len(match) > 0 {
			githubRepo.Color = match[0]
		}
		// repair before crawler no url
		githubRepo.Url = e.ChildAttr("h2 a", "href")
		// Kiểm tra và thêm "https://github.com/" nếu cần
		if !strings.HasPrefix(githubRepo.Url, "https://github.com/") {
			githubRepo.Url = "https://github.com" + githubRepo.Url
		}

		githubRepo.Lang = e.ChildText("span[itemprop=programmingLanguage]")

		e.ForEach(".mt-2 a", func(index int, el *colly.HTMLElement) {
			if index == 0 {
				githubRepo.Stars = strings.TrimSpace(el.Text)
			} else if index == 1 {
				githubRepo.Fork = strings.TrimSpace(el.Text)
			}
		})

		e.ForEach(".mt-2 .float-sm-right", func(index int, el *colly.HTMLElement) {
			githubRepo.StarsToday = strings.TrimSpace(el.Text)
		})

		var buildBy []string
		e.ForEach(".mt-2 span a img", func(index int, el *colly.HTMLElement) {
			avatarContributor := el.Attr("src")
			buildBy = append(buildBy, avatarContributor)
		})

		githubRepo.BuildBy = strings.Join(buildBy, ",")

		repos = append(repos, githubRepo)
	})

	c.OnScraped(func(r *colly.Response) {
		queue := NewJobQueue(runtime.NumCPU())
		queue.Start()
		defer queue.Stop()

		for _, repo := range repos {
			queue.Submit(&RepoProcess{
				repo:       repo,
				githubRepo: githubRepo,
			})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://github.com/trending")
}

type RepoProcess struct {
	repo       model.GithubRepo
	githubRepo repository.GithubRepo
}

func (rp *RepoProcess) Process() {
	// select repo by name
	cacheRepo, err := rp.githubRepo.SelectRepoByName(context.Background(), rp.repo.Name)
	if err == banana.RepoNotFound {
		// khong tim thay repo - insert repo to database
		fmt.Println("Add: ", rp.repo.Name)
		_, err = rp.githubRepo.SaveRepo(context.Background(), rp.repo)
		if err != nil {
			log.Error(err)
		}
		return
	}

	// Neu tim thấy thì update
	if rp.repo.Stars != cacheRepo.Stars ||
		rp.repo.StarsToday != cacheRepo.StarsToday ||
		rp.repo.Fork != cacheRepo.Fork {
		fmt.Println("Updated: ", rp.repo.Name)
		rp.repo.UpdatedAt = time.Now()
		_, err = rp.githubRepo.UpdateRepo(context.Background(), rp.repo)
		if err != nil {
			log.Error(err)
		}
	}
}
