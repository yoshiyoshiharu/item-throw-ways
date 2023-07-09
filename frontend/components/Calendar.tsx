import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";

const Calendar = () => (
  <FullCalendar
    plugins={[dayGridPlugin]}
    initialEvents={[{ title: "initial event", start: new Date() }]}
  />
 );

export default Calendar;
