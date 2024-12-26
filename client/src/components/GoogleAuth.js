import { GoogleLogin } from 'react-google-login';

const GoogleAuth = () => {
  const clientId = process.env.REACT_APP_GOOGLE_CLIENT_ID;

  const onSuccess = (response) => {
    console.log('Успешный вход:', response);
    // Отправляем токен на ваш бэкенд
    fetch(`${process.env.REACT_APP_API_URL}/auth/google`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ token: response.tokenId }),
    })
      .then((res) => res.json())
      .then((data) => {
        console.log('Ответ от сервера:', data);
      });
  };

  const onFailure = (error) => {
    console.error('Ошибка входа:', error);
  };

  return (
    <GoogleLogin
      clientId={clientId}
      buttonText="Войти через Google"
      onSuccess={onSuccess}
      onFailure={onFailure}
      cookiePolicy={'single_host_origin'}
    />
  );
};

export default GoogleAuth;