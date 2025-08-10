import './style.css'

// AccountButton.tsx
interface AccountButtonProps {
    accountName: string;
    balance: number;
    color: string; // Optional prop to change the background color of the button. Default is 'bg-blue-500'
}

const AccountButton = ({ accountName, balance, color }: AccountButtonProps) => {
    return (
        <div className={`${color} text-red rounded-xl flex flex-col items-start justify-center`}>

            <span className="text-sm font-medium">{accountName}</span>
            <span className="text-lg font-bold">
                ₹{balance.toLocaleString("en-IN", { minimumFractionDigits: 2 })}
            </span>
        </div>
    );
};

export default AccountButton;
