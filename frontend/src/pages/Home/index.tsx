import Footer from "@layout/Footer"
import Header from "@layout/Header"
import underConstructionImage from "@assets/underConstuction.gif"

const Home  =  () => {
    return(
     <>
     <Header/>
    <h2>Site under construction</h2>
    <img src={underConstructionImage} width="100%" alt="Under Construction"/>
    <Footer/>
    </>
    )
}


export default Home 