import EventEmitter from "eventemitter3";
import { useEffect } from "react";
import { Tracing } from "trace_events";

export class DataPublisher {
  private controller: DataController;
  private datatype: string;

  constructor(controller: DataController, datatype: string) {
    this.datatype = datatype;
    this.controller = controller;
    this.controller.Register(this.datatype, this);
  }

  Publish(data: any) {
    console.debug(`publishing to ${this.datatype}`);
    this.controller.Publish(this.datatype, data);
  }
}

export class DataSubscriber {
  private controller: DataController;
  private handler: Function;

  public OnChange: Function;

  constructor() {}

  Subscribe(controller: DataController, datatypes: string[]) {
    this.controller = controller;
  }

  Publish(datatype: string, data: any) {
    console.debug(`got data for to ${datatype}`);
    this.handler!(datatype, data);
    this.OnChange!(datatype, data);
  }
}

export class DataController {
  private data: Map<string, Array<DataSubscriber>>;
  private subscribers: Map<string, Array<DataSubscriber>>;
  private publishers: Map<string, Array<DataPublisher>>;

  constructor() {
    this.subscribers = new Map<string, Array<DataSubscriber>>();
    this.publishers = new Map<string, Array<DataPublisher>>();
  }

  private appendArray(
    map: Map<string, Array<any>>,
    datatype: string,
    publisher: DataPublisher | DataSubscriber
  ) {
    if (!map.has(datatype)) map.set(datatype, []);

    map.get(datatype)?.push(publisher);
  }

  Register(datatype: string, publisher: DataPublisher) {
    console.debug(`registration request for ${datatype}`);
    this.appendArray(this.publishers, datatype, publisher);
  }

  Subscribe(datatype: string, subscriber: DataSubscriber) {
    console.debug(`subscription request for ${datatype}`);
    this.appendArray(this.subscribers, datatype, subscriber);
  }

  Publish(datatype: string, data: any) {
    console.debug(
      `publishing request for ${datatype} to ${
        this.subscribers.get(datatype)?.length
      } subs`
    );
    this.subscribers.get(datatype)?.forEach((sub) => {
      sub.Publish(datatype, data);
    });
  }
}

const emitter = new EventEmitter();

export function UseSub(event, callback) {
  const unsubscribe = () => {
    emitter.off(event, callback);
  };

  useEffect(() => {
    emitter.on(event, callback);
    return unsubscribe;
  }, []);

  return unsubscribe;
}

export const usePub = () => {
  return (event, data) => {
    emitter.emit(event, data);
  };
};
