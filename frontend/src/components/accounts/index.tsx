// AccountContainer.tsx
import AccountButton from "./AccountButton";

const accountData = [
  { accountName: "Cash", balance: 1000, color: "bg-red-500" },
  { accountName: "Test", balance: 22, color: "bg-orange-500" },
  { accountName: "Jhu", balance: 20, color: "bg-blue-500" },
];

const AccountContainer = () => {
  return (
    <div className="">
      <h2 className="text-2xl font-bold mb-6 ">Accounts</h2>

      <div className="grid grid-cols-2 gap-4">
        {accountData.map((account) => (
          <AccountButton
            key={account.accountName}
            accountName={account.accountName}
            balance={account.balance}
            color={account.color}
          />
        ))}

        {/* Add Account Button */}
        <button className="flex flex-col justify-center items-center p-6 border-2 border-dashed rounded-xl text-gray-500 hover:text-blue-600 hover:border-blue-600 transition">
          <span className="text-lg font-medium">Add Account</span>
          <span className="text-2xl font-bold">＋</span>
        </button>
      </div>
    </div>
  );
};

export default AccountContainer;
