export const apiHost = import.meta.env.VITE_API_HOST;
export const gatewayKey = import.meta.env.VITE_GATEWAY_PUBLIC_KEY;

export async function uploadMedia(file, clientEmail, price, userID) {
    const formData = new FormData();
    formData.append('media', file);
    formData.append('email', clientEmail);
    formData.append('price', price);
    formData.append('user_id', userID);

    const response = await fetch(apiHost+'/upload', {
        method: 'POST',
        body: formData,
        credentials: 'include'
    });

    if (!response.ok) {
        throw new Error('Failed to upload file');
    }

    return response.json();
}

export async function logOut() {
    try {
        const response = await fetch(apiHost+'/logout', {
            method: 'POST',
            credentials: 'include'
        });
        if (response.ok) {
            window.location.href= '/';
        } else {
            alert('Failed to logout');
        }
    } catch (error) {
        console.error('Error logging out', error)
        alert('Error logging out.')
    }
}

export async function downloadMedia(mediaId) {
    try {
        const response = await fetch(apiHost+`/download/${mediaId}`, {
            method: 'POST',
            credentials: 'include',
            headers: {
                'Content-Type': 'application/json',
            },
            body : JSON.stringify({ mediaId })
        });

        if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
        }

        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);
        const a = document.createElement('a');
        a.href = url;
        a.download = `media_${mediaId}`;
        document.body.appendChild(a);
        a.click();
        a.remove();
        window.URL.revokeObjectURL(url);
    } catch(error) {
        console.error("Download failed:", error);
    }
}

export async function fetchMediaFiles() {
    const response = await fetch(apiHost+'/media', {
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
    });
    if (!response.ok) {
        throw new Error('Failed to fetch media files');
    }
    const data = await response.json();
    return data.media;
}

export async function checkSession() {
    const response = await fetch(apiHost+'/check', {
        credentials: 'include',
    });
    if (!response.ok) {
        return false;
    } else {
        const sessionData = await response.json();
        return sessionData.role === 'admin';
    }
}

export function formatDate(isoDate) {
    const date = new Date(isoDate);
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString(undefined, options);
}

