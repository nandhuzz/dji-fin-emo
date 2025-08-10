import djiFinLogo from "@assets/djifin.png"
import './style.css';
import { useNavigate } from "react-router-dom";
const Header = () => {
    const navigate = useNavigate();
    const handleClick = () => {
        navigate("/");
    }
    return (
        <div className="flex items-center h-[50px]" >
            <img src={djiFinLogo} className="w-[40px] mr-[2px]" alt="Dji Fin Logo" onClick={handleClick}/>
            <h1 className="text-2xl font-bold bg-clip-text text-transparent gold-silver-glow" onClick={handleClick}>
                ji Fin
            </h1>
        </div>
    )
}

export default Header