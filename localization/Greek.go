package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Greek] = map[string]string{
		"#admin":                              "διαχειριστής",
		"#admin.allow-role":                   "επιτρέπουν-ρόλο",
		"#look":                               "κοίτα",
		"#move":                               "κίνηση",
		"#review":                             "ανασκόπηση",
		"#admin.allow-role.value":             "αξία",
		"#.subject":                           "στόχος",
		"#look.info":                          "πληροφορίες",
		"#look.reviews":                       "κριτικές",
		"#admin.allow-role.Description":       "Εμφάνιση βαθμολογιών ως ρόλοι. Προεπιλογή: Λάθος",
		"#look.Description":                   "δείτε κάτι",
		"#move.Description":                   "Μετακίνηση κριτικής σε αυτό το κανάλι",
		"#review.Description":                 "Κριτική χρήστη",
		"#admin.allow-role.value.Description": "καθορισμένη τιμή",
		"#.subject.Description":               "Επιλέξτε στόχο",
		"#look.info.Description":              "Προβολή πληροφοριών χρήστη",
		"#look.reviews.Description":           "Δείτε τις κριτικές που έχει λάβει το θέμα",
		"#allow-role.NeedPermissions":         "Το ρομπότ δεν διαθέτει δικαιώματα - Διαχείριση ρόλων",
		"#allow-role.InProgress":              "εργασία σε εξέλιξη",
		"#allow-role.proc.Title":              "Τροποποίηση Επιλογών",
		"#allow-role.proc.Description":        "Επιτρέψτε ρόλους",
		"#allow-role.proc.InProgress":         "Ενέργεια",
		"#allow-role.proc.Done":               "πλήρης",
		"#allow-role.Keep":                    "Δεν άλλαξαν ρυθμίσεις",
		"#look.info.IsNone":                   "Δεν υπάρχουν κριτικές",
		"#look.reviews.IsNone":                "Δεν υπάρχουν κριτικές",
		"#look.reviews.menu.Title":            "Κριτικές για %s",
		"#look.reviews.menu.Page":             "%d σελίδα",
		"#move.IsNone":                        "Δεν υπάρχουν κριτικές για τον στόχο",
		"#move.Move":                          "ελα εδω",
		"#review.SelfReview":                  "Δεν μπορείς να αναθεωρήσεις τον εαυτό σου",
		"#review.modal.Title":                 "Έλεγχος %s",
		"#review.lable.Score":                 "σκορ",
		"#review.lable.Title":                 "τίτλος",
		"#review.lable.Content":               "λεπτομέρεια",
		"$review.IsEdited":                    "Αυτή η κριτική έχει τροποποιηθεί",
		"$review.NoAuthor":                    "Δεν είναι δυνατή η επαναφορά των διαγραμμένων κριτικών επειδή ο συγγραφέας δεν υπάρχει εδώ",
		"$review.DM":                          "Μια νέα κριτική έχει γραφτεί",
    }
}
