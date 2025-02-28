import Table from "../Table/Table";

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
