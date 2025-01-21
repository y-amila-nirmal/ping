import React from "react";

function App() {
    const [networkData, setNetworkData] = React.useState({
        bytesSent: 0,
        bytesRecv: 0,
        packetsSent: 0,
        packetsRecv: 0,
    });

    React.useEffect(() => {
        (window as any).runtime.EventsOn("networkData", (data: any) => {
            setNetworkData(data);
        })

        return () => {
            (window as any).runtime.EventsOff("networkData");
        }
    }, [networkData]);

    return (
        <div>
            <p>Network Usage</p>
            <p>Bytes Sent: {networkData.bytesSent}</p>
            <p>Bytes Recv: {networkData.bytesRecv}</p>
            <p>Packets Sent: {networkData.packetsSent}</p>
            <p>Packets Recv: {networkData.packetsRecv}</p>
        </div>
    );
}

export default App;