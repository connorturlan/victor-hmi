import { useState } from "react";
import Table from "../Table/Table";
import styles from "./TimesheetTable.module.scss";

const timecodeMappings = { USAC25: ["Monday", "Thursday", "Saturday"] };
const weekdays = [
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
  "Sunday",
];

function TimesheetTable({ tableData, onUpdate }) {
  const handleUpdate = (data) => {
    const newTable = data.map((row) => {
      const timecode = row.at(0);
      const validDays = timecodeMappings[timecode];
      console.log(timecode, validDays);
      if (!validDays) return row;

      row = Array.apply(null, Array(row.length)).map(() => -1);
      row[0] = timecode;
      for (const day of validDays) {
        console.log(day, weekdays.indexOf(day));
        row[weekdays.indexOf(day) + 1] = 0;
      }
      return row;
    });
    onUpdate(newTable);
  };

  return (
    <Table
      headings={["Timecode", ...weekdays]}
      tableData={tableData}
      onUpdate={handleUpdate}
    ></Table>
  );
}

export default TimesheetTable;
