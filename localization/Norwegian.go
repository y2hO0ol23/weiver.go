package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Norwegian] = map[string]string{
		"#admin":                              "admin",
		"#admin.allow-role":                   "tillate-rolle",
		"#look":                               "se",
		"#move":                               "bevege-seg",
		"#review":                             "anmeldelse",
		"#admin.allow-role.value":             "verdi",
		"#.subject":                           "mål",
		"#look.info":                          "informasjon",
		"#look.review-list":                   "gjennomgå-liste",
		"#admin.allow-role.Description":       "Vis poeng som roller. Standard: False",
		"#look.Description":                   "se noe",
		"#move.Description":                   "Flytt anmeldelse til denne kanalen",
		"#review.Description":                 "Brukeranmeldelse",
		"#admin.allow-role.value.Description": "angi verdi",
		"#.subject.Description":               "Velg mål",
		"#look.info.Description":              "Se brukerinformasjon",
		"#look.review-list.Description":       "Se en liste over anmeldelser mottatt av brukere",
		"#allow-role.NeedPermissions":         "Bot mangler tillatelser - Administrer roller",
		"#allow-role.InProgress":              "arbeid pågår",
		"#allow-role.proc.Title":              "Endre alternativer",
		"#allow-role.proc.Description":        "Tillat roller",
		"#allow-role.proc.InProgress":         "Går videre",
		"#allow-role.proc.Done":               "fullstendig",
		"#allow-role.Keep":                    "Ingen innstillinger endret",
		"#look.info.IsNone":                   "Ingen anmeldelser",
		"#look.review-list.IsNone":            "Det er ingen anmeldelser",
		"#look.review-list.menu.Title":        "Anmeldelser for %s",
		"#look.review-list.menu.Page":         "%d side",
		"#move.IsNone":                        "Det er ingen anmeldelser skrevet for målet",
		"#move.Move":                          "Flytt her",
		"#review.SelfReview":                  "Du kan ikke anmelde deg selv",
		"#review.modal.Title":                 "Gjennomgå %s",
		"#review.lable.Score":                 "score",
		"#review.lable.Title":                 "tittel",
		"#review.lable.Content":               "detalj",
		"$review.IsEdited":                    "Denne anmeldelsen er redigert",
		"$review.NoAuthor":                    "Slettede anmeldelser kan ikke gjenopprettes fordi forfatteren ikke eksisterer her",
		"$review.DM":                          "En ny anmeldelse er skrevet",
    }
}