import { useEffect, useState } from "react";
import { json } from "stream/consumers";

interface user {
    map(arg0: (item: any) => JSX.Element): import('react').ReactNode;
    username: string;
    last_played: EpochTimeStamp;
}

function Users() {
    const [users, setUsers] = useState<user[] | undefined>()

    useEffect(() => {
        fetch('http://' + window.location.hostname + ':8080/users')
            .then((response) => response.json())
            .then((json) => setUsers(json.data));
    }, [])

    return <div>
        <h2>Current Users are : </h2>
        {users ? <ul>
            {users.map((item: any) => {
                return <li key={item.username}>{item.username} - {item.last_played}</li>
            })}
        </ul> : <p>Loading ...</p>}
    </div>
}

export default Users;