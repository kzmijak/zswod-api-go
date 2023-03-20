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
	}

	pages := []CustomPageHead{
		{ section: "O Szkole", icon: "Enlist", title: "Rekrutacja", },
		{ section: "O Szkole", icon: "Documents", title: "Szkolne dokumenty", },
		{ section: "O Szkole", icon: "Chronicle", title: "Kronika szkoły", },
		{ section: "O Szkole", icon: "Library", title: "Biblioteka", },
		{ section: "O Szkole", icon: "Contests", title: "Konkursy", },
		{ section: "O Szkole", icon: "Projects", title: "Projekty", },
		{ section: "O Szkole", icon: "Talentowisko", title: "Talentowisko", link: "https://www.talentowisko.pl/podstawowa/szkola/szkola-podstawowa-im-jana-pawla-ii-w-orlowie-drewnianym" },
		{ section: "Dla Rodziców", icon: "Meeting", title: "Harmonogram zebrań", },
		{ section: "Dla Rodziców", icon: "Consult", title: "Konsultacje dla Rodziców", },
		{ section: "Dla Rodziców", icon: "Counsil", title: "Rada Rodziców", },
		{ section: "Dla Rodziców", icon: "Staff", title: "Grono pedagogiczne", },
		{ section: "Dla Rodziców", icon: "Lessons", title: "Kalendarz roku szkolnego", },
		{ section: "Dla Rodziców", icon: "Help", title: "Pomoc psychologiczno-pedagogiczna", },
		{ section: "Dla Ucznia", icon: "Lessons", title: "Tygodniowy Plan Zajęć", },
		{ section: "Dla Ucznia", icon: "Grades", title: "Dziennik elektroniczny", },
		{ section: "Dla Ucznia", icon: "StudentCounsil", title: "Samorząd Uczniowski", },
		{ section: "Dla Ucznia", icon: "Extracurricular", title: "Zajęcia pozalekcyjne", },
		{ section: "Dla Ucznia", icon: "Pedagogue", title: "Pedagog szkolny", },
		{ section: "Dla Ucznia", icon: "Volunteer", title: "Szkolne Koło Wolontariatu", },
		{ section: "Dla Ucznia", icon: "Mentors", title: "Wychowawcy", },
	}

	bulk := make([]*ent.CustomPageCreate, len(pages))

	for i, page := range pages {
		titleSanitized := sanitizer.SanitizeString(page.title)

		bulk[i] = Client.CustomPage.Create().
			SetSection(page.section).
			SetIconId(page.icon).
			SetTitle(page.title).
			SetTitleNormalized(titleSanitized).
			SetContent("")

		if (page.link != "") {
			bulk[i] = bulk[i].
				SetIsExternal(true).
				SetLink(page.link)
		}
	}

	_, err = Client.CustomPage.CreateBulk(bulk...).Save(ctx)

	return err
}
