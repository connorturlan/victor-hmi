import { useState } from "react";
import {
  DataController,
  DataPublisher,
  DataSubscriber,
} from "../../core/datasources";
import styles from "./DataContainer.module.scss";
import DataSub from "../../components/DataSub/DataSub";
import DataPub from "../../components/DataPub/DataPub";

function DataContainer(props) {
  const controller = new DataController();

  return (
    <div>
      <DataPub controller={controller} />
      <DataSub controller={controller} />
    </div>
  );
}

export default DataContainer;
