import HeaderBar from './components/headerbar';
import MainContent from './components/maincontent';
import Footer from './components/footer';
import './App.css';

function App() {
  return (
    <div className="app">
      <HeaderBar />
      <div className="container-fluid">
        <MainContent />
      </div>
      <Footer />
    </div>
  );
}

export default App;
