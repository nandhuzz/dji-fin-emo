// AccountContainer.tsx

import AccountButton from "./AccountButton";

const accountData = [
  { accountName: "Cash", balance: 1000, color: "bg-red-500" },
  { accountName: "test", balance: 22, color: "bg-orange-500" },
  { accountName: "jhu", balance: 20, color: "bg-blue-500" },
];

const AccountContainer = () => {
  return (
    <div className="p-4">
      <h2 className="text-xl font-bold mb-4">List of accounts</h2>
      <div className="">
        {accountData.map((account) => (
          <AccountButton
            key={account.accountName}
            accountName={account.accountName}
            balance={account.balance}
            color={account.color}
          />
        ))}
        {/* Add Account Button */}
        <div className="">
          <span className="">ADD ACCOUNT</span>
          <span className="">＋</span>
        </div>
      </div>
    </div>
  );
};

export default AccountContainer;
