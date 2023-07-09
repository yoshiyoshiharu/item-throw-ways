import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import { AreaCollectionDate } from "@/entity/area_collection_date";

export default function Calendar({ areaCollectionDates }: { areaCollectionDates: AreaCollectionDate[] }) {
  // const events = areaCollectionDates.map((collectionDate: AreaCollectionDate) => {
  //   return {
  //     title: 'ゴミの日',
  //     start: '2023-07-04',
  //   }
  // });
  // 
  const events = [
    { title: 'event 1', start: '2023-07-04' },
    { title: 'event 2', start: '2023-07-05' }
  ]
  console.log(events)

  return(
    <FullCalendar
      plugins={[dayGridPlugin]}
      initialEvents={events}
      contentHeight="auto"
      locale="ja"
    />
  )
}
