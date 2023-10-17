package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.German] = map[string]string{
		"#admin":                              "administrator",
		"#admin.allow-role":                   "rolle-zulassen",
		"#look":                               "sehen",
		"#move":                               "bewegen",
		"#review":                             "rezension",
		"#admin.allow-role.value":             "wert",
		"#.subject":                           "ziel",
		"#look.info":                          "information",
		"#look.reviews":                       "bewertungen",
		"#admin.allow-role.Description":       "Punktestände als Rollen anzeigen. Standard: Falsch",
		"#look.Description":                   "etwas besichtigen",
		"#move.Description":                   "Rezension auf diesen Kanal verschieben",
		"#review.Description":                 "Nutzerbewertung",
		"#admin.allow-role.value.Description": "Wert einstellen",
		"#.subject.Description":               "Ziel auswählen",
		"#look.info.Description":              "Benutzerinformationen anzeigen",
		"#look.reviews.Description":           "Sehen Sie sich die Bewertungen an, die das Thema erhalten hat",
		"#admin.allow-role.NeedPermissions":   "Dem Bot fehlen Berechtigungen – Rollen verwalten",
		"#admin.allow-role.InProgress":        "in Arbeit",
		"#admin.allow-role.proc.Title":        "Optionen ändern",
		"#admin.allow-role.proc.Description":  "Rollen zulassen",
		"#admin.allow-role.proc.InProgress":   "Weiter",
		"#admin.allow-role.proc.Done":         "vollständig",
		"#admin.allow-role.Keep":              "Keine Einstellungen geändert",
		"#look.info.IsNone":                   "Keine Bewertungen",
		"#look.reviews.IsNone":                "Es gibt keine Bewertungen",
		"#look.reviews.menu.Title":            "Bewertungen für %s",
		"#look.reviews.menu.Page":             "%d Seite",
		"#move.IsNone":                        "Für das Ziel liegen keine Rezensionen vor",
		"#move.Move":                          "Komm her",
		"#review.SelfReview":                  "Sie können sich nicht selbst bewerten",
		"#review.modal.Title":                 "Überprüfen Sie %s",
		"#review.lable.Score":                 "Punktzahl",
		"#review.lable.Title":                 "Titel",
		"#review.lable.Content":               "Detail",
		"$review.IsEdited":                    "Diese Rezension wurde bearbeitet",
		"$review.NoAuthor":                    "Gelöschte Rezensionen können nicht wiederhergestellt werden, da der Autor hier nicht existiert",
		"$review.DM":                          "Eine neue Rezension wurde verfasst",
	}
}
