import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import { AreaCollectionDate } from "@/entity/area_collection_date";

export default function Calendar({ areaCollectionDates }: { areaCollectionDates: AreaCollectionDate[] }) {
  const backgroundColor = (kind: string) => {
    switch (kind) {
      case "可燃ごみ":
        return "#DC2626";
      case "不燃ごみ":
        return "#2563EB";
      case "資源":
        return "#15803D";
      default:
        return "gray";
    }
  }

  const events = areaCollectionDates.map((areaCollectionDate: AreaCollectionDate) => {
    return {
      title: areaCollectionDate.kind.name,
      start: areaCollectionDate.date,
      backgroundColor: backgroundColor(areaCollectionDate.kind.name),
      borderColor: backgroundColor(areaCollectionDate.kind.name)
    }
  });

  return(
    <FullCalendar
      plugins={[dayGridPlugin]}
      events={events}
      contentHeight="auto"
      locale="ja"
      headerToolbar={{
        left:   'title',
        center: '',
        right:  ''
      }}
    />
  )
}
