package database

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
	"github.com/kzmijak/zswod_api_go/utils/sanitizer"
)

var Client *ent.Client; 


func InitDatabase(cfg DatabaseConfig, ctx context.Context) error  {

	client, err := ent.Open(cfg.DriverName, cfg.DSN )
	if err != nil {
		return ErrConnectionFailed
	}

	Client = client;
	
	if err := Client.Schema.Create(ctx); err != nil {
		return ErrSchemaCreationFail
	}


	if cfg.Env == "dev" {
		if err := seedAdmin(ctx); err != nil {
			return err
		}
	}


	if err := seedPaths(ctx); err != nil {
		return err
	}
	
	return nil;
}

func seedAdmin(ctx context.Context) error {
	adminsCount, err := Client.User.Query().Where(user.RoleEQ(user.RoleAdmin)).Count(ctx)
	if err != nil {
		return ErrCouldNotQuery
	}

	if adminsCount > 0 {
		return nil
	}

	adminPassword, _ := encryption.HashString("root")

	err = Client.User.Create().SetID(uuid.UUID{}).SetEmail("root@sporlowd.pl").SetPassword(adminPassword).SetRole(user.RoleAdmin).Exec(ctx)

	return err
}

func seedPaths(ctx context.Context) error {
	pagesCount, err := Client.CustomPage.Query().Count(ctx)
	if err != nil {
		return ErrCouldNotQuery
	}

	if pagesCount > 0 {
		return nil
	}

	type CustomPageHead struct {
		createTime time.Time
		icon string
		section string
		title string
		link string
		order int
	}

	pages := []CustomPageHead{
		{ createTime: time.Now().Add(0), section: "O Szkole", icon: "Enlist", title: "Rekrutacja" },
		{ createTime: time.Now().Add(1), section: "O Szkole", icon: "Documents", title: "Szkolne dokumenty", },
		{ createTime: time.Now().Add(2), section: "O Szkole", icon: "Chronicle", title: "Kronika szkoły", },
		{ createTime: time.Now().Add(3), section: "O Szkole", icon: "Library", title: "Biblioteka", },
		{ createTime: time.Now().Add(4), section: "O Szkole", icon: "Contests", title: "Konkursy", },
		{ createTime: time.Now().Add(5), section: "O Szkole", icon: "Projects", title: "Projekty", },
		{ createTime: time.Now().Add(6), section: "O Szkole", icon: "Talentowisko", title: "Talentowisko", link: "https://www.talentowisko.pl/podstawowa/szkola/szkola-podstawowa-im-jana-pawla-ii-w-orlowie-drewnianym" },
		{ createTime: time.Now().Add(7), section: "Dla Rodziców", icon: "Meeting", title: "Harmonogram zebrań", },
		{ createTime: time.Now().Add(8), section: "Dla Rodziców", icon: "Consult", title: "Konsultacje dla Rodziców", },
		{ createTime: time.Now().Add(9), section: "Dla Rodziców", icon: "Counsil", title: "Rada Rodziców", },
		{ createTime: time.Now().Add(10), section: "Dla Rodziców", icon: "Staff", title: "Grono pedagogiczne", },
		{ createTime: time.Now().Add(11), section: "Dla Rodziców", icon: "Lessons", title: "Kalendarz roku szkolnego", },
		{ createTime: time.Now().Add(12), section: "Dla Rodziców", icon: "Help", title: "Pomoc psychologiczno-pedagogiczna", },
		{ createTime: time.Now().Add(13), section: "Dla Ucznia", icon: "Lessons", title: "Tygodniowy Plan Zajęć", },
		{ createTime: time.Now().Add(14), section: "Dla Ucznia", icon: "Grades", title: "Dziennik elektroniczny", },
		{ createTime: time.Now().Add(15), section: "Dla Ucznia", icon: "StudentCounsil", title: "Samorząd Uczniowski", },
		{ createTime: time.Now().Add(16), section: "Dla Ucznia", icon: "Extracurricular", title: "Zajęcia pozalekcyjne", },
		{ createTime: time.Now().Add(17), section: "Dla Ucznia", icon: "Pedagogue", title: "Pedagog szkolny", },
		{ createTime: time.Now().Add(18), section: "Dla Ucznia", icon: "Volunteer", title: "Szkolne Koło Wolontariatu", },
		{ createTime: time.Now().Add(19), section: "Dla Ucznia", icon: "Mentors", title: "Wychowawcy", },
	}

	bulk := make([]*ent.CustomPageCreate, len(pages))

	for i, page := range pages {
		sectionSanitized := sanitizer.SanitizeString(page.section)
		titleSanitized := sanitizer.SanitizeString(page.title)
		url := sectionSanitized + "/" + titleSanitized;

		bulk[i] = Client.CustomPage.Create().
			SetCreateTime(page.createTime).
			SetSection(page.section).
			SetIconId(page.icon).
			SetTitle(page.title).
			SetURL(url).
			SetContent("")

		if (page.link != "") {
			bulk[i] = bulk[i].
				SetLink(page.link)
		}
	}

	_, err = Client.CustomPage.CreateBulk(bulk...).Save(ctx)

	return err
}
