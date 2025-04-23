import { lazy, useState } from "react";
import styles from "./DataContainer.module.scss";
import DataSub from "../../components/DataSub/DataSub";
import DataPub from "../../components/DataPub/DataPub";

function DataContainer(props) {
  return lazy(
    <div>
      <DataPub controller={controller} />
      <DataSub controller={controller} />
    </div>
  );
}

export default DataContainer;
