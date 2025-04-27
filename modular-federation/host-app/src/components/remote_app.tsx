import { lazy, Suspense } from "react";
import {
  __federation_method_getRemote,
  __federation_method_setRemote,
  // @ts-ignore
} from "__federation__";

// const RemoteAsset = lazy(() => import("remote_app_1/Asset"));
const RemoteAsset = lazy(() => {
  // values like { 'http://localhost:9000/assets/remoteEntry.js', 'remoteA', './RemoteARoot' }
  const url = "http://localhost:4173/assets/remoteEntry.js";
  const name = "remote_app_1";
  const module = "./Asset";

  __federation_method_setRemote(name, {
    url: () => Promise.resolve(url),
    format: "esm",
    from: "vite",
  });

  return __federation_method_getRemote(name, module);
});

export const Asset: React.FC = () => {
  return (
    <Suspense fallback="loading">
      <RemoteAsset></RemoteAsset>
    </Suspense>
  );
};
