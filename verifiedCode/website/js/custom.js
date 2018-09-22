var app = angular.module('fileUpload', []);
app.controller('ctrlFileUpload', ['$scope', 'uploadFactory', function($scope, uploadFactory) {
    $scope.form = [];
    $scope.UploadMyFiles = function() {
        console.log("test");
    };

    $scope.clickUpload = function(){
        console.log("a");
        angular.element('#upload').trigger('click');	    
    };

    $scope.onFileSelect = function(element){
        console.log("b");
        $scope.fileinprocess = true;
        $scope.no_files = false;
        $scope.currentFile = element.files[0];
        var reader = new FileReader();
        reader.onload = function (event) {
            $scope.image_source = event.target.result
            $scope.$apply(function ($scope) {
                $scope.files = element.files;
            });
        };
        reader.readAsDataURL(element.files[0]);
        setTimeout(function() {
            $scope.form.image = $scope.files[0];
            uploadFactory.uploadfile_tos3($scope.form.image, $scope.form, 1).then(function(response){
                $scope.filename = (response.data).replace('\n\n','');
                response = (response.data).replace('\n\n','');
                $scope.fileExists = true;
                $scope.fileinprocess = false;
            });
        }, 500);
    };
}
]);
