import { useState } from 'react';

const tabs = [
    { id: "files", title: "Files" },
    { id: "settings", title: "Settings" },
    { id: "info", title: "Info" },
];

const Tray = () => {
    const [activeTab, setActiveTab] = useState("");
    return (
        <div className="tray d-flex flex-column">
            <ul className="nav nav-tabs flex-row">
                {tabs.map(tab => (
                    <button className={`nav-link flex-fill ${activeTab === tab.id ? 'active' : ''}`} onClick={() => setActiveTab(activeTab === tab.id ? "" : tab.id)}>
                        {tab.title}
                    </button>
                ))}
            </ul>
            {
                activeTab !== "" &&
                <div className="tab-content flex-grow-1 p-3 border">
                    {activeTab === "files" && <div>Files content</div>}
                    {activeTab === "settings" && <div>Settings content</div>}
                    {activeTab === "info" && <div>Info content</div>}
                </div>
            }
        </div>
    );
};

export default Tray;