module.exports = {
    default: {
        require: ['src/features/step_definitions/**/*.js'],
        format: ['progress', 'json:reports/cucumber_report.json'],
        paths: ['src/features/**/*.feature'],
    }
};
