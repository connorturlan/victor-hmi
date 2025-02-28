import { useEffect, useState } from "react";
import styles from "./Table.module.scss";

function Table({ headings, data, onUpdate }) {
  const [tableData, setRows] = useState(data);

  const updateRow = (rowIndex, cellIndex, value) => {
    const newTable = tableData.slice();
    newTable[rowIndex][cellIndex] = value;
    setRows(newTable);
  };

  const addRow = () => {
    const newTable = tableData.slice();
    newTable.push(Array.apply(0, Array(headings.length)));
    setRows(newTable);
  };

  const removeRow = (rowIndex) => {
    console.log(rowIndex);
    const preTable = tableData.slice(0, rowIndex);
    const postTable = tableData.slice(rowIndex + 1);
    const newTable = preTable.concat(postTable);
    console.log(preTable, postTable, newTable);
    setRows(newTable);
  };

  useEffect(() => {
    console.log("updating table");
    onUpdate && onUpdate(tableData);
  }, [tableData]);

  return (
    <div className={styles.Table}>
      <table>
        <thead>
          <tr>
            {headings.map((heading) => {
              return <th key={`heading${heading}`}>{heading}</th>;
            })}
          </tr>
        </thead>
        <tbody>
          {tableData.map((row, rowIndex) => {
            return (
              <tr key={`row${rowIndex}`}>
                {row.map((cell, cellIndex) => {
                  return (
                    <td key={`cell${rowIndex},${cellIndex}`}>
                      <input
                        placeholder="0"
                        value={cell}
                        onChange={(event) =>
                          updateRow(rowIndex, cellIndex, event.target.value)
                        }
                      ></input>
                    </td>
                  );
                })}
                <td>
                  <button onClick={() => removeRow(rowIndex)}>
                    Remove Row
                  </button>
                </td>
              </tr>
            );
          })}
          <tr>
            <td>
              <button onClick={addRow}>Add Row</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}

export default Table;
