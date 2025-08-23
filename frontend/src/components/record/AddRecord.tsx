import { useEffect, useState } from "react";

const AddRecord = () => {
    const [record, setRecord] = useState({
        transactionType: "",
        amount: "",
        account: "",
        category: { parent: "", sub: "" },
        description: "",
        labels: "",
        payee: "",
        date: "",
        time: "",
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setRecord((prev) => ({ ...prev, [name]: value }));
    };

    const setRecordType = (type) => {
        setRecord((prev) => ({ ...prev, transactionType: type }));
    };

    useEffect(() => {
        const t = new Date();

        const yyyy = t.getFullYear();
        const mm = String(t.getMonth() + 1).padStart(2, "0");
        const dd = String(t.getDate()).padStart(2, "0");

        const today = `${yyyy}-${mm}-${dd}`; // for input type="date"
        const time = t.toTimeString().slice(0, 5); // HH:MM

        setRecord((prev) => ({
            ...prev,
            date: today,
            time: time,
        }));
    }, []);

    

    const handleSubmit = () => {
        console.log("Saving record:", record);
        // Call API here
    };

    return (
        <div className="max-w-xl mx-auto p-6 shadow-md rounded-2xl space-y-6">
            {/* Record Type */}
            <div className="flex justify-around">
                {["income", "expense", "transfer"].map((type) => (
                    <button
                        key={type}
                        onClick={() => setRecordType(type)}
                        className={`px-4 py-2 rounded-xl font-semibold ${record.transactionType === type
                            ? "bg-blue-600 text-white"
                            : "bg-gray-200 text-gray-700"
                            }`}
                    >
                        {type.toUpperCase()}
                    </button>
                ))}
            </div>

            {/* Amount */}
            <div>
                <label className="block text-gray-700 font-semibold mb-1">Amount</label>
                <input
                    type="number"
                    name="amount"
                    value={record.amount}
                    onChange={handleChange}
                    placeholder="0"
                    className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                />
            </div>

            {/* Account & Category */}
            <div className="grid grid-cols-2 gap-4">
                <div>
                    <label className="block text-gray-700 font-semibold mb-1">Account</label>
                    <input
                        type="text"
                        name="account"
                        value={record.account}
                        onChange={handleChange}
                        placeholder="Cash / Bank"
                        className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                    />
                </div>
                <div>
                    <label className="block text-gray-700 font-semibold mb-1">Category</label>
                    <input
                        type="text"
                        name="category.parent"
                        value={record.category.parent}
                        onChange={(e) =>
                            setRecord((prev) => ({
                                ...prev,
                                category: { ...prev.category, parent: e.target.value },
                            }))
                        }
                        placeholder="Shopping"
                        className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                    />
                </div>
            </div>

            {/* Description */}
            <div>
                <label className="block text-gray-700 font-semibold mb-1">Description</label>
                <textarea
                    name="description"
                    value={record.description}
                    onChange={handleChange}
                    placeholder="Enter description"
                    className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                />
            </div>

            {/* Labels */}
            <div>
                <label className="block text-gray-700 font-semibold mb-1">Labels</label>
                <input
                    type="text"
                    name="labels"
                    value={record.labels}
                    onChange={handleChange}
                    placeholder="Comma separated labels"
                    className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                />
            </div>

            {/* Payee */}
            <div>
                <label className="block text-gray-700 font-semibold mb-1">Payee</label>
                <input
                    type="text"
                    name="payee"
                    value={record.payee}
                    onChange={handleChange}
                    placeholder="Enter payee name"
                    className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                />
            </div>

            {/* Date & Time */}
            <div className="grid grid-cols-2 gap-4">
                <div>
                    <label className="block text-gray-700 font-semibold mb-1">Date</label>
                    <input
                        type="date"
                        name="date"
                        value={record.date}
                        onChange={handleChange}
                        className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                    />
                </div>
                <div>
                    <label className="block text-gray-700 font-semibold mb-1">Time</label>
                    <input
                        type="time"
                        name="time"
                        value={record.time}
                        onChange={handleChange}
                        className="w-full px-4 py-2 border rounded-xl focus:ring-2 focus:ring-blue-500"
                    />
                </div>
            </div>

            {/* Save Button */}
            <button
                onClick={handleSubmit}
                className="w-full py-3 rounded-xl bg-blue-600 text-white font-semibold hover:bg-blue-700 transition"
            >
                Save Record
            </button>
        </div>
    );
};

export {AddRecord};
