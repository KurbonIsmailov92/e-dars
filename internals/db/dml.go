package db

const (
	DBGetTeacherID = `SELECT DISTINCT u.id
						FROM journal_notes jn
								 JOIN schedules s ON jn.schedule_id = s.id
								 JOIN classes c ON s.class_id = c.id
								 JOIN class_users cu ON c.id = cu.class_id
								 JOIN users u ON cu.user_id = u.id
						WHERE c.id = ?`

	DBGetChildJournalNotesByDates = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name
									FROM journal_notes jn
											 JOIN schedules s ON jn.schedule_id = s.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN groups g ON s.group_id = g.id
											 JOIN users stud ON g.id = stud.group_id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON teach.id = cu.user_id
											 LEFT JOIN marks m ON jn.mark_id = m.id
									WHERE stud.parent_id = ?
									  AND jn.date::date BETWEEN ? AND ?`

	DBGetJournalNotesByID = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name
									FROM journal_notes jn
											 JOIN schedules s ON jn.schedule_id = s.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN groups g ON s.group_id = g.id
											 JOIN users stud ON g.id = stud.group_id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON teach.id = cu.user_id
											 LEFT JOIN marks m ON jn.mark_id = m.id
									WHERE jn.id = ?`

	DBGetOwnJournalNotesByDates = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name
									FROM journal_notes jn
											 JOIN schedules s ON jn.schedule_id = s.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN groups g ON s.group_id = g.id
											 JOIN users stud ON g.id = stud.group_id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON teach.id = cu.user_id
											 LEFT JOIN marks m ON jn.mark_id = m.id
									WHERE stud.id = ?
									  AND jn.date::date BETWEEN ? AND ?`

	DBGetJournalNotesByTeacherAndDates = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name
									FROM journal_notes jn
											 JOIN schedules s ON jn.schedule_id = s.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN groups g ON s.group_id = g.id
											 JOIN users stud ON g.id = stud.group_id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON teach.id = cu.user_id
											 LEFT JOIN marks m ON jn.mark_id = m.id
									WHERE teach.id = ?
									  AND jn.date::date BETWEEN ? AND ?`
)
