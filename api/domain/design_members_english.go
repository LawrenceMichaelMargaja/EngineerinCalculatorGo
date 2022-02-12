package domain

type DesignMembersEnglish struct {
	ID                            int    `json:"id" db:"id"`
	DesignMembersEnglishName       string `json:"design_members_english_name" db:"design_members_english_name"`
	DesignMembersEnglishTotalDepth string `json:"design_members_english_total_depth" db:"design_members_english_total_depth"`
	DesignMembersEnglishWeight     string `json:"design_members_english_weight" db:"design_members_english_weight"`
}
