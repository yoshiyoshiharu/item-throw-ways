import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import { AreaCollectionDate } from "@/entity/area_collection_date";

export default function Calendar({ areaCollectionDates }: { areaCollectionDates: AreaCollectionDate[] }) {
  const events = areaCollectionDates.map((areaCollectionDate: AreaCollectionDate) => {
    return {
      title: areaCollectionDate.kind,
      start: areaCollectionDate.date
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
