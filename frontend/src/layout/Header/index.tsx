import djiFinLogo from "@assets/djifin.png"
import './style.css';
const Header = () => {
    return (
        <div className="flex items-center h-[50px] glow-ray">
            <img src={djiFinLogo} className="w-[40px] mr-[2px]" alt="Dji Fin Logo" />
            <h1 className="text-2xl font-bold bg-clip-text text-transparent gold-silver-glow">
                ji Fin
            </h1>
        </div>
    )
}

export default Header