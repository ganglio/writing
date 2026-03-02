import HeaderBar from './components/headerbar';
import MainContent from './components/maincontent';
import Footer from './components/footer';

import GlobalContext from './context/global';

import './App.css';

function App() {

    return (
        <GlobalContext.Provider>
            <div className="app d-flex flex-column min-vh-100">
                <div className="row">
                    <HeaderBar />
                </div>
                <div className="row flex-grow-1">
                    <MainContent />
                </div>
                <div className="row">
                    <Footer />
                </div>
            </div>
        </GlobalContext.Provider>
    );
}

export default App;
