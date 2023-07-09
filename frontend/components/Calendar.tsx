import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import { AreaCollectionDate } from "@/entity/area_collection_date";
import { useState } from "react";

export default function Calendar({ areaCollectionDates }: { areaCollectionDates: AreaCollectionDate[] }) {
  // const events = areaCollectionDates.map((collectionDate: AreaCollectionDate) => {
  //   return {
  //     title: collectionDate.kind,
  //     start: collectionDate.date,
  //   }
  // })

  return(
    <FullCalendar
      plugins={[dayGridPlugin]}
      // initialEvents={events}
      contentHeight="auto"
      locale="ja"
    />
  )
}
