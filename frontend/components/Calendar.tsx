import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";

const Calendar = () => (
  <FullCalendar
    plugins={[dayGridPlugin]}
    initialEvents={[
      {
        start: new Date()
      },
      {
        title: "不燃ごみ",
        start: '2023-07-18',
        backgroundColor: 'blue',
      },
      {
        title: "可燃ごみ",
        start: '2023-07-20',
        backgroundColor: 'red',
      }
    ]}
    contentHeight="auto"
    locale="ja"
  />
 );

export default Calendar;
