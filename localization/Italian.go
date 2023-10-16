package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Italian] = map[string]string{
		"#admin":                              "amministratore",
		"#admin.allow-role":                   "consentire-il-ruolo",
		"#look":                               "aspetto",
		"#move":                               "mossa",
		"#review":                             "revisione",
		"#admin.allow-role.value":             "valore",
		"#.subject":                           "bersaglio",
		"#look.info":                          "informazione",
		"#look.review-list":                   "elenco-di-revisione",
		"#admin.allow-role.Description":       "Visualizza i punteggi come ruoli. Predefinito: falso",
		"#look.Description":                   "vedere qualcosa",
		"#move.Description":                   "Sposta la recensione su questo canale",
		"#review.Description":                 "Recensione dell'utente",
		"#admin.allow-role.value.Description": "valore impostato",
		"#.subject.Description":               "Seleziona destinazione",
		"#look.info.Description":              "Visualizza le informazioni dell'utente",
		"#look.review-list.Description":       "Visualizza un elenco delle recensioni ricevute dagli utenti",
		"#allow-role.NeedPermissions":         "Il bot non dispone delle autorizzazioni: Gestisci ruoli",
		"#allow-role.InProgress":              "lavori in corso",
		"#allow-role.proc.Title":              "Modifica opzioni",
		"#allow-role.proc.Description":        "Consenti ruoli",
		"#allow-role.proc.InProgress":         "Procedere",
		"#allow-role.proc.Done":               "completare",
		"#allow-role.Keep":                    "Nessuna impostazione modificata",
		"#look.info.IsNone":                   "Nessuna recensione",
		"#look.review-list.IsNone":            "Non ci sono recensioni",
		"#look.review-list.menu.Title":        "Recensioni per %s",
		"#look.review-list.menu.Page":         "%d pagina",
		"#move.IsNone":                        "Non ci sono recensioni scritte per il target",
		"#move.Move":                          "Vieni qui",
		"#review.SelfReview":                  "Non puoi rivedere te stesso",
		"#review.modal.Title":                 "Recensioni",
		"#review.lable.Score":                 "punto",
		"#review.lable.Title":                 "titolo",
		"#review.lable.Content":               "dettaglio",
		"$review.IsEdited":                    "Questa recensione è stata modificata",
		"$review.NoAuthor":                    "Le recensioni cancellate non possono essere ripristinate perché l'autore non esiste qui",
		"$review.DM":                          "È stata scritta una nuova recensione",
    }
}
