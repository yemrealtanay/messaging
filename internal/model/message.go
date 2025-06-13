package model

type Message struct {
	ID        int64          `json:"id" db:"id"`
	ToPhone   string         `json:"to_phone" db:"to_phone"`
	Content   string         `json:"content" db:"content"`
	IsSent    bool           `json:"is_sent" db:"is_sent"`
	SentAt    *FormattedTime `json:"sent_at" db:"sent_at"`
	MessageID *string        `json:"message_id" db:"message_id"`
	CreatedAt *FormattedTime `json:"created_at,omitempty" db:"created_at"`
}

//NOT: Gerçek bir senaryoda production'da message modelinin içinde to_phone
//tutulmasını doğru bulmuyorum. Aslında yapı şöyle olmalı;
//contact vardır burada contact'ın isim soyisim ve diğer ana bilgilerini tutarız.
//contact_comm_infos vardır burada contact_id ile communication infolar bağlarız.
//hatta belki communication info type gibi bir email/phone durumunu da tutmalıyız.
//mesaj comm_info_id tutmalı. to_phone tutmak bana doğru gelmiyor.
//ama mülakat taskının basitliğini korumak için bu yola başvurdum.

//ayrıca created_at eklemek istedim. Çünkü takip edilebilirlik ve sıralamada işe yarar diye düşünüyorum.
//belki debug için updated_at ve soft delete amacıyla deleted_at de ekleyebilirdim. ancak yine
//taskı basit tutmak istedim.

//laraveldeki cast yapısına benzer bir tarih formatlama yapısı kullanmayı seviyorum.
//bu yüzden burada formatted_time modelini kullandım.
