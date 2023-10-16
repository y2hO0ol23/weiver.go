package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Croatian] = map[string]string{
		"#admin":                              "admin",
		"#admin.allow-role":                   "dopustiti-ulogu",
		"#look":                               "izgled",
		"#move":                               "potez",
		"#review":                             "pregled",
		"#admin.allow-role.value":             "vrijednost",
		"#.subject":                           "cilj",
		"#look.info":                          "informacija",
		"#look.review-list":                   "popis-pregleda",
		"#admin.allow-role.Description":       "Prikaži rezultate kao uloge. Zadano: False",
		"#look.Description":                   "vidjeti nešto",
		"#move.Description":                   "Premjesti recenziju na ovaj kanal",
		"#review.Description":                 "Pregled korisnika",
		"#admin.allow-role.value.Description": "postavljena vrijednost",
		"#.subject.Description":               "Odaberite cilj",
		"#look.info.Description":              "Pregledajte podatke o korisniku",
		"#look.review-list.Description":       "Pogledajte popis recenzija koje su primili korisnici",
		"#allow-role.NeedPermissions":         "Botu nedostaju dopuštenja - Upravljanje ulogama",
		"#allow-role.InProgress":              "radovi u tijeku",
		"#allow-role.proc.Title":              "Modificiraj opcije",
		"#allow-role.proc.Description":        "Dopusti uloge",
		"#allow-role.proc.InProgress":         "Postupak",
		"#allow-role.proc.Done":               "potpuna",
		"#allow-role.Keep":                    "Nema promjena postavki",
		"#look.info.IsNone":                   "Nema recenzija",
		"#look.review-list.IsNone":            "Nema recenzija",
		"#look.review-list.menu.Title":        "Recenzije za %s",
		"#look.review-list.menu.Page":         "%d stranica",
		"#move.IsNone":                        "Nema napisanih recenzija za cilj",
		"#move.Move":                          "preseli ovamo",
		"#review.SelfReview":                  "Ne možete se pregledati",
		"#review.modal.Title":                 "Pregledajte %s",
		"#review.lable.Score":                 "postići",
		"#review.lable.Title":                 "titula",
		"#review.lable.Content":               "detalj",
		"$review.IsEdited":                    "Ova recenzija je uređena",
		"$review.NoAuthor":                    "Izbrisane recenzije nije moguće vratiti jer autor ne postoji ovdje",
		"$review.DM":                          "Napisana je nova recenzija",
    }
}
