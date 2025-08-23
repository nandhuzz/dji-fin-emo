import Footer from "@layout/Footer"
import Header from "@layout/Header"
import AccountContainer from "@components/accounts"
import Transactions from "@components/transactions"
import "./style.css"
import AddRecord from "@/components/record"

const Home = () => {
    return (
        <div className="homeContainer">
            <Header />
            <AccountContainer />
            <Transactions />
            <AddRecord />
            <Footer />
        </div>
    )
}


export default Home