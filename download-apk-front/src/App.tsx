import { QRCode } from 'react-qrcode-logo';
import './App.css';
import { useEffect, useState } from 'react';

function App() {
  const [urlString, setUrlString] = useState(
    'https://github.com/asma12a/challenge-s6/releases/download/v.1.0/app-release.apk'
  );

  const getUrl = async () => {
    const url =
      'https://api.github.com/repos/asma12a/challenge-s6/releases/latest';

    const response = await fetch(url);
    const data = await response.json();
    if (response.status !== 200) {
      console.error('Latest release not found');
      return;
    }
    const apkDownloadUrl = data.assets.find((asset: { name: string }) =>
      asset.name.endsWith('.apk')
    ).browser_download_url;
    setUrlString(apkDownloadUrl);
  };

  useEffect(() => {
    getUrl();
  }, []);

  return (
    <>
      <h1>Télécharger l'application Squad Go</h1>
      <QRCode value={urlString} size={500} qrStyle='dots' eyeRadius={10} />
    </>
  );
}

export default App;
