export const getTest = () => {
    return new Promise(resolve => {
        setTimeout(() => {
            resolve();
        }, 2000);
    })
};