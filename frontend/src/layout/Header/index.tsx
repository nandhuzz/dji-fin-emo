import djiFinLogo from "@assets/djifin.png"

const Header = () => {
    return (
        <div className="flex items-center gap-4 p-4">
            <img src={djiFinLogo} width="100px" alt="Dji Fin Logo" />
            <h1 className="text-2xl font-bold">Dji Fin</h1>
        </div>
    )
}

export default Header