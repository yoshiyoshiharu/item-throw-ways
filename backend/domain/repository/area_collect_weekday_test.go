package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yoshiyoshiharu/item-throw-ways/infrastructure/entity"
)

func TestAreaCollectWeekdayRepository_FindByAreaId(t *testing.T) {
	gormDB, mock := newMockDB()

	repo := NewAreaCollectWeekdayRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "area_id", "weekday", "lap"}).
		AddRow(1, 1, 3, 0).
		AddRow(2, 1, 5, 0)

	mock.ExpectQuery("SELECT `area_collect_weekdays`.`id`,`area_collect_weekdays`.`area_id`,`area_collect_weekdays`.`kind_id`,`area_collect_weekdays`.`weekday`,`area_collect_weekdays`.`lap`,`Kind`.`id` AS `Kind__id`,`Kind`.`name` AS `Kind__name`,`Area`.`id` AS `Area__id`,`Area`.`name` AS `Area__name` FROM `area_collect_weekdays` LEFT JOIN `kinds` `Kind` ON `area_collect_weekdays`.`kind_id` = `Kind`.`id` LEFT JOIN `areas` `Area` ON `area_collect_weekdays`.`area_id` = `Area`.`id` WHERE area_id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	areaCollectWeekdays := repo.FindByAreaId(1)

	t.Run("[正常系] FindByIdは指定したIDのエリアを返すこと", func(t *testing.T) {
		assert.Equal(t, 1, areaCollectWeekdays[0].ID)
		assert.Equal(t, 2, areaCollectWeekdays[1].ID)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestAreaCollectWeekdayRepository_DeleteAndInsertAll(t *testing.T) {
	gormDB, mock := newMockDB()

	repo := NewAreaCollectWeekdayRepository(gormDB)

	area := entity.NewArea(1, "Area 1")
	kind := entity.NewKind(1, "Kind 1")

	areaCollectWeekday1 := entity.NewAreaCollectWeekday(area, kind, 3, 0)
	areaCollectWeekday2 := entity.NewAreaCollectWeekday(area, kind, 5, 0)

	areaCollectWeekdays := []*entity.AreaCollectWeekday{
		areaCollectWeekday1,
		areaCollectWeekday2,
	}

	t.Run("[正常系] areasとarea_collect_weekdaysを全消去し、areasとarea_collect_weekdaysを挿入すること", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM areas").WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectExec("DELETE FROM area_collect_weekdays").WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO `areas` (`name`,`id`) VALUES (?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
			WithArgs(area.Name, area.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO `area_collect_weekdays` (`area_id`,`kind_id`,`weekday`,`lap`) VALUES (?,?,?,?),(?,?,?,?)")).
			WithArgs(areaCollectWeekday1.AreaId, areaCollectWeekday1.KindId, areaCollectWeekday1.Weekday, areaCollectWeekday1.Lap, areaCollectWeekday2.AreaId, areaCollectWeekday2.KindId, areaCollectWeekday2.Weekday, areaCollectWeekday2.Lap).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.DeleteAndInsertAll(areaCollectWeekdays)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("[異常系] トランザクションでロールバックした場合はエラーを返すこと", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM areas").WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectExec("DELETE FROM area_collect_weekdays").WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO `areas` (`name`,`id`) VALUES (?,?) ON DUPLICATE KEY UPDATE `id`=`id`")).
			WithArgs(area.Name, area.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec(
			regexp.QuoteMeta("INSERT INTO `area_collect_weekdays` (`area_id`,`kind_id`,`weekday`,`lap`) VALUES (?,?,?,?),(?,?,?,?)")).
			WithArgs(areaCollectWeekday1.AreaId, areaCollectWeekday1.KindId, areaCollectWeekday1.Weekday, areaCollectWeekday1.Lap, areaCollectWeekday2.AreaId, areaCollectWeekday2.KindId, areaCollectWeekday2.Weekday, areaCollectWeekday2.Lap).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectRollback()

		err := repo.DeleteAndInsertAll(areaCollectWeekdays)

		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}
