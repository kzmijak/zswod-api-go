package database

import (
	"context"

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
		icon string
		section string
		title string
		link string
		order int
	}

	pages := []CustomPageHead{
		{ order: 0, section: "O Szkole", icon: "Enlist", title: "Rekrutacja" },
		{ order: 1, section: "O Szkole", icon: "Documents", title: "Szkolne dokumenty", },
		{ order: 2, section: "O Szkole", icon: "Chronicle", title: "Kronika szkoły", },
		{ order: 3, section: "O Szkole", icon: "Library", title: "Biblioteka", },
		{ order: 4, section: "O Szkole", icon: "Contests", title: "Konkursy", },
		{ order: 5, section: "O Szkole", icon: "Projects", title: "Projekty", },
		{ order: 6, section: "O Szkole", icon: "Talentowisko", title: "Talentowisko", link: "https://www.talentowisko.pl/podstawowa/szkola/szkola-podstawowa-im-jana-pawla-ii-w-orlowie-drewnianym" },
		{ order: 7, section: "Dla Rodziców", icon: "Meeting", title: "Harmonogram zebrań", },
		{ order: 8, section: "Dla Rodziców", icon: "Consult", title: "Konsultacje dla Rodziców", },
		{ order: 9, section: "Dla Rodziców", icon: "Counsil", title: "Rada Rodziców", },
		{ order: 10, section: "Dla Rodziców", icon: "Staff", title: "Grono pedagogiczne", },
		{ order: 11, section: "Dla Rodziców", icon: "Lessons", title: "Kalendarz roku szkolnego", },
		{ order: 12, section: "Dla Rodziców", icon: "Help", title: "Pomoc psychologiczno-pedagogiczna", },
		{ order: 13, section: "Dla Ucznia", icon: "Lessons", title: "Tygodniowy Plan Zajęć", },
		{ order: 14, section: "Dla Ucznia", icon: "Grades", title: "Dziennik elektroniczny", },
		{ order: 15, section: "Dla Ucznia", icon: "StudentCounsil", title: "Samorząd Uczniowski", },
		{ order: 16, section: "Dla Ucznia", icon: "Extracurricular", title: "Zajęcia pozalekcyjne", },
		{ order: 17, section: "Dla Ucznia", icon: "Pedagogue", title: "Pedagog szkolny", },
		{ order: 18, section: "Dla Ucznia", icon: "Volunteer", title: "Szkolne Koło Wolontariatu", },
		{ order: 19, section: "Dla Ucznia", icon: "Mentors", title: "Wychowawcy", },
	}

	bulk := make([]*ent.CustomPageCreate, len(pages))

	for i, page := range pages {
		sectionSanitized := sanitizer.SanitizeString(page.section)
		titleSanitized := sanitizer.SanitizeString(page.title)
		url := sectionSanitized + "/" + titleSanitized;

		bulk[i] = Client.CustomPage.Create().
			SetOrder(page.order).
			SetSection(page.section).
			SetIconId(page.icon).
			SetTitle(page.title).
			SetTitleNormalized(url).
			SetContent("")

		if (page.link != "") {
			bulk[i] = bulk[i].
				SetLink(page.link)
		}
	}

	_, err = Client.CustomPage.CreateBulk(bulk...).Save(ctx)

	return err
}
