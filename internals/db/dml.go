package db

const (
	GetTeacherIDDB = `SELECT DISTINCT u.id
						FROM journal_notes jn
								 JOIN schedules s ON jn.schedule_id = s.id
								 JOIN classes c ON s.class_id = c.id
								 JOIN class_users cu ON c.id = cu.class_id
								 JOIN users u ON cu.user_id = u.id
						WHERE c.id = ?`

	GetChildJournalNotesByDatesDB = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name,
										   s.is_exam 		  as is_exam
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

	GetJournalNotesByIDDB = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name,
										   s.is_exam 		  as is_exam
									FROM journal_notes jn
											 JOIN schedules s ON jn.schedule_id = s.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN groups g ON s.group_id = g.id
											 JOIN users stud ON g.id = stud.group_id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON teach.id = cu.user_id
											 LEFT JOIN marks m ON jn.mark_id = m.id
									WHERE jn.id = ?`

	GetOwnJournalNotesByDatesDB = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name,
										   s.is_exam 		  as is_exam
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

	GetJournalNotesByTeacherAndDatesDB = `SELECT jn.date,
										   g.code_name        as "group",
										   stud.full_name     as student_name,
										   c.name             as class,
										   c.classroom_number as class_room,
										   m.code             as mark,
										   teach.full_name    as teacher_name,
										   s.is_exam 		  as is_exam
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

	GetTeacherScheduleByDateDB = `SELECT s.planned_date     as date,
										   c.name             as class,
										   c.classroom_number as class_room,
										   g.code_name        as "group",
										   u.full_name        as teacher,
										   s.is_exam 		  as is_exam
									FROM schedules s
											 JOIN groups g ON s.group_id = g.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users u ON cu.user_id = u.id
									WHERE u.id = ?
									  AND s.planned_date::DATE BETWEEN ? and ?`

	GetStudentScheduleByDateDB = `SELECT s.planned_date     as date,
										   c.name             as class,
										   c.classroom_number as class_room,
										   g.code_name        as "group",
										   teach.full_name    as teacher,
										   s.is_exam 		  as is_exam
									FROM schedules s
											 JOIN groups g ON s.group_id = g.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON cu.user_id = teach.id
									WHERE g.id = (SELECT DISTINCT stud.group_id
												  FROM users stud
												  WHERE stud.id = ?)
									  AND s.planned_date::DATE BETWEEN ? and ?`

	GetParentScheduleByDateDB = `SELECT s.planned_date     as date,
										   c.name             as class,
										   c.classroom_number as class_room,
										   g.code_name        as "group",
										   teach.full_name    as teacher,
										   s.is_exam 		  as is_exam
									FROM schedules s
											 JOIN groups g ON s.group_id = g.id
											 JOIN classes c ON s.class_id = c.id
											 JOIN class_users cu ON c.id = cu.class_id
											 JOIN users teach ON cu.user_id = teach.id
									WHERE g.id in (SELECT stud.group_id
												  FROM users stud
												  WHERE stud.parent_id = ?)
									  AND s.planned_date::DATE BETWEEN ? and ?`
)
