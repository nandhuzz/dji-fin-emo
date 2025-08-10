import Footer from "@layout/Footer"
import Header from "@layout/Header"
import AccountContainer from "@components/accounts"
import Transactions from "@components/transactions"
import "./style.css"

const Home = () => {
    return (
        <div className="homeContainer">
            <Header />
            <AccountContainer />
            <Transactions />
            <Footer />
        </div>
    )
}


export default Home