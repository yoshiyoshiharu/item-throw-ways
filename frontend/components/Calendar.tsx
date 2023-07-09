import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import { AreaCollectionDate } from "@/entity/area_collection_date";

export default function Calendar({ areaCollectionDates }: { areaCollectionDates: AreaCollectionDate[] }) {
  const backgroundColor = (kind: string) => {
    switch (kind) {
      case "可燃ごみ":
        return "red";
      case "不燃ごみ":
        return "blue";
      case "資源":
        return "green";
      default:
        return "gray";
    }
  }

  const events = areaCollectionDates.map((areaCollectionDate: AreaCollectionDate) => {
    return {
      title: areaCollectionDate.kind,
      start: areaCollectionDate.date,
      backgroundColor: backgroundColor(areaCollectionDate.kind)
    }
  });

  return(
    <FullCalendar
      plugins={[dayGridPlugin]}
      events={events}
      contentHeight="auto"
      locale="ja"
    />
  )
}
