package database

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (d *Database) runSchema() {
	createColumns := `
	CREATE TABLE IF NOT EXISTS columns (
		id VARCHAR(50) PRIMARY KEY,
		title VARCHAR(100) NOT NULL
	);`

	createTasks := `
	CREATE TABLE IF NOT EXISTS tasks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title TEXT NOT NULL,
		description TEXT,
		assignee VARCHAR(100),
		due_date DATE,
		created_at DATE DEFAULT CURRENT_DATE,
		priority VARCHAR(10) CHECK (priority IN ('low','medium','high')),
		is_blocked BOOLEAN DEFAULT FALSE,
		column_id VARCHAR(50) NOT NULL REFERENCES columns(id) ON DELETE CASCADE,
		project VARCHAR(100),
		completion_percentage INT CHECK (completion_percentage BETWEEN 0 AND 100)
	);`

	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		username VARCHAR(100) PRIMARY KEY,
		password TEXT NOT NULL
	);`

	alterAssignee := `
	DO $$
	BEGIN
		IF EXISTS (
			SELECT 1 FROM information_schema.columns 
			WHERE table_name='tasks' AND column_name='assignee'
		) AND NOT EXISTS (
			SELECT 1 FROM information_schema.table_constraints 
			WHERE table_name='tasks' AND constraint_type='FOREIGN KEY' AND constraint_name='fk_tasks_assignee'
		) THEN
			ALTER TABLE tasks
			ADD CONSTRAINT fk_tasks_assignee FOREIGN KEY (assignee)
			REFERENCES users(username)
			ON DELETE SET NULL;
		END IF;
	END;
	$$;`

	_, err := d.DB.Exec(createColumns)
	if err != nil {
		log.Fatalf("❌ Error creando tabla columns: %v", err)
	}
	log.Println("✅ Tabla 'columns' creada/verificada")

	_, err = d.DB.Exec(createTasks)
	if err != nil {
		log.Fatalf("❌ Error creando tabla tasks: %v", err)
	}
	log.Println("✅ Tabla 'tasks' creada/verificada")

	_, err = d.DB.Exec(createUsers)
	if err != nil {
		log.Fatalf("❌ Error creando tabla users: %v", err)
	}
	log.Println("✅ Tabla 'users' creada/verificada")

	_, err = d.DB.Exec(alterAssignee)
	if err != nil {
		log.Fatalf("❌ Error añadiendo clave foránea a assignee: %v", err)
	}
	log.Println("🔗 Clave foránea 'assignee → users.username' verificada")

	// 👤 Usuario por defecto
	const username = "admin"
	const password = "admin01"

	var exists bool
	err = d.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM users WHERE username = $1)`, username).Scan(&exists)
	if err != nil {
		log.Fatalf("❌ Error verificando existencia del usuario admin: %v", err)
	}

	if !exists {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("❌ Error generando hash bcrypt: %v", err)
		}

		_, err = d.DB.Exec(`INSERT INTO users (username, password) VALUES ($1, $2)`, username, string(hashed))
		if err != nil {
			log.Fatalf("❌ Error insertando usuario admin: %v", err)
		}

		log.Println("🔐 Usuario inicial 'admin' creado")
	} else {
		log.Println("🔐 Usuario 'admin' ya existe")
	}

	log.Println("🧱 Esquema creado/verificado con éxito")
}
