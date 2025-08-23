import { useState, useEffect } from "react";

// Fake API (replace later with your real backend)
const fetchTransactions = async (page = 1) => {
  const data = [
    {
      id: "tans_1",
      type: "income",
      amount: 3000,
      account: "Cash",
      date: "2025-08-22",
      time: "10:30",
    },
    {
      id: "tans_2",
      type: "income",
      amount: 2000,
      account: "Bank",
      date: "2025-08-22",
      time: "14:45",
    },
    {
      id: "tans_3",
      type: "expense",
      amount: 4000,
      account: "Card",
      date: "2025-08-23",
      time: "09:10",
    },
  ];

  return data.slice(0, page * 2); // simulate pagination
};

const Transactions = () => {
  const [transactions, setTransactions] = useState([]);
  const [page, setPage] = useState(1);

  useEffect(() => {
    const loadTransactions = async () => {
      const result = await fetchTransactions(page);
      setTransactions(result);
    };
    loadTransactions();
  }, [page]);

  const loadMore = () => {
    setPage((prev) => prev + 1);
  };

  return (
    <div className="">
      <h2 className="text-xl font-bold mb-4">Transactions</h2>

      <ul className="divide-y divide-gray-200 border rounded-lg ">
        {transactions.map((item) => (
          <li key={item.id} className="flex justify-between p-4 hover:bg-gray-50">
            <div>
              <p
                className={`font-semibold ${
                  item.type === "income" ? "text-green-600" : "text-red-600"
                }`}
              >
                ₹{item.amount}
              </p>
              <p className="text-sm text-gray-600 capitalize">{item.type}</p>
            </div>
            <div className="text-right text-sm text-gray-500">
              <p>
                {item.date} {item.time}
              </p>
              <p>{item.account}</p>
            </div>
          </li>
        ))}
      </ul>

      <div className="mt-4 flex justify-center">
        <button
          onClick={loadMore}
          className="px-6 py-2 rounded-lg bg-blue-600 text-white font-semibold hover:bg-blue-700 transition"
        >
          Load More
        </button>
      </div>
    </div>
  );
};

export { Transactions };
