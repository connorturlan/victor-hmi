export const loadRemoteModule = async (appName: string, moduleName: string) => {
  const MAX_RETRIES = 3;
  let retries = 0;

  while (retries < MAX_RETRIES) {
    try {
      return await import(`${appName}/${moduleName}`);
    } catch (err) {
      retries++;
      await new Promise((resolve) => setTimeout(resolve, 1000));
    }
  }
  throw new Error(`Failed to load ${moduleName}`);
};
