import { ElementType, lazy, Suspense, useState } from "react";
import {
  __federation_method_getRemote,
  __federation_method_setRemote,
  // @ts-ignore
} from "__federation__";

interface DynamicRemoteAssetProps {
  url: string;
  name: string;
  module: string;
}

const DynamicRemoteAsset:
  | React.LazyExoticComponent<React.ComponentType<any>>
  | any = (props: DynamicRemoteAssetProps) => {
  const { url, name, module } = props;
  if (!url || !name || !module) return <></>;

  const Component = lazy(() => {
    __federation_method_setRemote(name, {
      url: () => Promise.resolve(url),
      format: "esm",
      from: "vite",
    });

    return __federation_method_getRemote(name, module);
  });

  const loader = <h1>loading bruh</h1>;

  return (
    <Suspense fallback={loader}>
      <Component />
    </Suspense>
  );
};

export const DynamicAsset: ElementType = () => {
  const [url, setUrl] = useState("");
  const [name, setName] = useState("");

  return (
    <>
      <button
        onClick={() => {
          setUrl("http://localhost:4173/assets/remoteEntry.js");
          setName("remote_app_1");
        }}
      >
        app 1
      </button>
      <button
        onClick={() => {
          setUrl("http://localhost:4174/assets/remoteEntry.js");
          setName("remote_app_2");
        }}
      >
        app 2
      </button>
      <DynamicRemoteAsset url={url} name={name} module="./Asset" />
    </>
  );
};
