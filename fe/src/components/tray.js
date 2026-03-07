import { useState, useEffect, useContext } from 'react';

import GlobalContext from '../context/global';

const Tray = () => {
    const [activeTab, setActiveTab] = useState("");
    const [tabs, setTabs] = useState([]);

    const { state, dispatch } = useContext(GlobalContext.Context);


    useEffect(() => {
        // Simulate fetching tabs from an API
        setTimeout(() => {
            setTabs([
                { id: "files", icon: "fa-regular fa-folder" },
                { id: "settings", icon: "fa-solid fa-gear" },
                { id: "info", icon: "fa-solid fa-circle-info" }
            ]);
        }, 1000);
    }, []);

    return (
        <div className="tray d-flex flex-column">
            { tabs.length === 0 ?
                <div className="p-2 text-center">Loading tabs...</div> :
                <ul className="nav nav-tabs flex-row">
                    {
                        tabs.map(tab => (
                            <button key={tab.id} className={`nav-link flex-fill ${activeTab === tab.id ? 'active' : ''}`} onClick={() => setActiveTab(activeTab === tab.id ? "" : tab.id)}>
                                {
                                    activeTab === tab.id ? <i className="fa-solid fa-circle-xmark" /> : <i className={tab.icon} />
                                }
                            </button>
                        ))
                    }
                    <button className="nav-link">
                        <i className="fa-solid fa-plus" onClick={()=>dispatch({
                            type: GlobalContext.Actions.TOGGLE_MODAL,
                            payload: "createautomation"
                        })}/>
                    </button>
                </ul>
            }
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