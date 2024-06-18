package database

func (d *Database) CreateUserTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS Users (
            id INT PRIMARY KEY AUTO_INCREMENT,
            full_name VARCHAR(255) NOT NULL,
            username VARCHAR(50) NOT NULL,
            password VARCHAR(255) NOT NULL,
            email VARCHAR(100) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            UNIQUE (username),
            UNIQUE (email)
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateConversationTable() error {
	query := `
        CREATE TABLE Conversations (
            id INT PRIMARY KEY AUTO_INCREMENT,
            user_id INT NOT NULL,
            status VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES Users(id)
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateMessageTable() error {
	query := `
        CREATE TABLE Messages (
            id INT PRIMARY KEY AUTO_INCREMENT,
            conversation_id INT NOT NULL,
            user_id INT NOT NULL,
            message TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (conversation_id) REFERENCES Conversations(id),
            FOREIGN KEY (user_id) REFERENCES Users(id)
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateIntentTable() error {
	query := `
        CREATE TABLE Intents (
            id INT PRIMARY KEY AUTO_INCREMENT,
            name VARCHAR(255) NOT NULL,
            description TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateEntityTable() error {
	query := `
        CREATE TABLE Entities (
            id INT PRIMARY KEY AUTO_INCREMENT,
            name VARCHAR(255) NOT NULL,
            value VARCHAR(255) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateResponseTable() error {
	query := `
        CREATE TABLE Responses (
            id INT PRIMARY KEY AUTO_INCREMENT,
            intent_id INT NOT NULL,
            message TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            FOREIGN KEY (intent_id) REFERENCES Intents(id)
        );
    `
	_, err := d.DB.Exec(query)
	return err
}

func (d *Database) CreateAllTables() error {
	if err := d.CreateUserTable(); err != nil {
		return err
	}
	if err := d.CreateConversationTable(); err != nil {
		return err
	}
	if err := d.CreateMessageTable(); err != nil {
		return err
	}
	if err := d.CreateIntentTable(); err != nil {
		return err
	}
	if err := d.CreateEntityTable(); err != nil {
		return err
	}
	if err := d.CreateResponseTable(); err != nil {
		return err
	}
	return nil
}
