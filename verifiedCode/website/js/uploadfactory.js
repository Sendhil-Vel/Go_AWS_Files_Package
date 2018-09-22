'use strict';

app.factory('uploadFactory', ['$rootScope', '$http', '$q', '$location',
    function($rootScope, $http, $q, $location) {
		var emailBaseURL = 'http://localhost:8080';
		return ({
			/* uploading the file to S3 */
            uploadfile_tos3: uploadfile_tos3
		});	
		function uploadfile_tos3(uploadHandler, form, UserID) {
			return $http({
				method: 'POST',
				url: emailBaseURL + '/uploadfile',
				processData: false,
				transformRequest: function (data) {
					var formData = new FormData();
					formData.append("file", uploadHandler);
					return formData;
				},
				data: form,
				headers: {
					'Content-Type': undefined
				}
			}).then(handleSuccess, handleError);
		}
		/* Process error response data here: */
        function handleError(response) {
            console.log(response);
            if (response.status === 401 && response.statusText === 'User not authorized') {
                $location.path('/');
            }
            return ($q.reject(response));
        }

        /* Process success data here: */
        function handleSuccess(response) {

            if (response.config.method === 'POST') {
                $rootScope.$emit(response.data.exitcode === 1 ? 'serverError' : 'serverSuccess', response.data.message);
            }
            return (response);
        }
	}
]);