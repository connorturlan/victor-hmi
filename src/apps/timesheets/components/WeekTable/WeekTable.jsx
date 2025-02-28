import Table from "../Table/Table";
import styles from "./WeekTable.module.scss";

function WeekTable({ weeks, onUpdate, onClick }) {
  return (
    <Table
      headings={["Week Ending"]}
      tableData={weeks}
      onUpdate={onUpdate}
      onClick={onClick}
    ></Table>
  );
}

export default WeekTable;
