import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";

const Calendar = () => (
  <FullCalendar
    plugins={[dayGridPlugin]}
    initialEvents={[{ start: new Date() }]}
    contentHeight="auto"
    locale="ja"
  />
 );

export default Calendar;
