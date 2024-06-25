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
        CREATE TABLE IF NOT EXISTS Conversations (
            id INT PRIMARY KEY AUTO_INCREMENT,
            user_id INT NOT NULL,
            title VARCHAR(255) NOT NULL,
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
        CREATE TABLE IF NOT EXISTS Messages (
            id INT PRIMARY KEY AUTO_INCREMENT,
            conversation_id INT NOT NULL,
            user_id INT NOT NULL,
            message TEXT NOT NULL,
            role VARCHAR(50) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (conversation_id) REFERENCES Conversations(id),
            FOREIGN KEY (user_id) REFERENCES Users(id)
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
	return nil
}
