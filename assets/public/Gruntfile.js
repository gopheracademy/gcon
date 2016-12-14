'use strict';
module.exports = function(grunt) {
    [
        'grunt-contrib-watch',
        'grunt-contrib-clean',
        'grunt-contrib-copy',
        'grunt-contrib-less'
    ].forEach(function(task) { grunt.loadNpmTasks(task); });

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),

        clean: {
            dist: {
                src: [
                    '!less/**.less',
                    '!less/lib/*.css',
                    'less/*.css',
                ]
            }
        },

        copy: {
            dist: {
                files: [{
                    expand: true,
                    dot: true,
                    cwd: '',
                    dest: '',
                    src: [
                        '*',
                        'less/**',
                    ],
                    filter: 'isFile'
                }]
            }
        },

        less: {
            development: {
                options: {
                    paths: ['less'],
                    ieCompat: false
                },
                files: {
                    'css/style.css': 'less/styles/style.less',
                    'css/custom-animations.css': 'less/animations.less'
                }
            },
            color: {
                options: {
                    paths: ['less'],
                    ieCompat: false
                },
                files: {
                    'css/style-blue-color.css': 'less/styles/style-blue-color.less',
                    'css/style-pink-color.css': 'less/styles/style-pink-color.less',
                    'css/style-green-color.css': 'less/styles/style-green-color.less',
                    'css/style-berry-color.css': 'less/styles/style-berry-color.less',
                    'css/style-orange-color.css': 'less/styles/style-orange-color.less'
                }
            },
            production: {
                options: {
                    compress: true,
                    yuicompress: false,
                    paths: ['less'],
                    ieCompat: false
                },
                files: {
                    'css/min/style.min.css': 'less/styles/style.less',
                    'css/min/custom-animations.min.css': 'less/animations.less',
                    'css/min/style-blue-color.min.css': 'less/styles/style-blue-color.less',
                    'css/min/style-pink-color.min.css': 'less/styles/style-pink-color.less',
                    'css/min/style-green-color.min.css': 'less/styles/style-green-color.less',
                    'css/min/style-berry-color.min.css': 'less/styles/style-berry-color.less',
                    'css/min/style-orange-color.min.css': 'less/styles/style-orange-color.less'
                }
            }
        },

        watch: {
            less: {
                files: ['**/*.less'],
                tasks: ['less:development']
            },
            livereload: {
                options: { livereload: true },
                files: ['**']
            }
        }
    });

    grunt.registerTask('develop', ['watch']);

    grunt.registerTask('production', ['clean', 'less:development', 'less:color', 'less:production', 'copy']);

    grunt.registerTask('default', ['clean', 'less:development', 'copy']);
};
